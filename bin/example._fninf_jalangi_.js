J$.iids = {"8":[23,9,23,30],"9":[5,28,5,29],"10":[23,15,23,25],"16":[23,9,23,30],"17":[5,28,5,29],"18":[23,9,23,25],"24":[26,9,26,19],"25":[5,22,5,30],"26":[26,9,26,19],"32":[36,9,36,27],"33":[5,1,5,32],"34":[27,13,27,18],"40":[36,9,36,27],"41":[5,1,5,32],"42":[36,9,36,21],"48":[48,9,48,15],"49":[5,1,5,32],"50":[36,25,36,27],"56":[54,5,54,20],"57":[11,14,11,22],"58":[47,9,47,14],"65":[11,14,11,22],"66":[48,9,48,15],"73":[11,14,11,22],"74":[54,5,54,20],"81":[12,14,12,22],"89":[12,14,12,22],"97":[12,14,12,22],"105":[14,9,14,15],"113":[14,9,14,15],"121":[14,9,14,15],"129":[15,9,15,14],"137":[15,9,15,14],"145":[15,9,15,14],"153":[20,2,20,40],"161":[20,2,20,41],"169":[21,2,21,63],"177":[21,2,21,64],"185":[23,9,23,10],"193":[23,15,23,21],"201":[23,24,23,25],"209":[23,29,23,30],"217":[24,15,24,38],"225":[24,15,24,38],"233":[24,9,24,39],"241":[26,9,26,10],"249":[26,13,26,19],"257":[27,13,27,14],"265":[27,17,27,18],"273":[27,13,27,18],"281":[27,9,27,19],"289":[28,9,28,21],"297":[28,22,28,30],"305":[28,9,28,31],"313":[28,9,28,31],"321":[19,1,30,2],"329":[19,1,30,2],"337":[33,2,33,41],"345":[33,2,33,42],"353":[34,2,34,70],"361":[34,2,34,71],"369":[36,9,36,10],"377":[36,15,36,21],"385":[36,26,36,27],"393":[37,9,37,21],"401":[37,22,37,29],"409":[37,9,37,30],"417":[37,9,37,31],"425":[39,9,39,21],"433":[39,22,39,29],"441":[39,9,39,30],"449":[39,9,39,31],"457":[32,1,41,2],"465":[32,1,41,2],"473":[44,2,44,41],"481":[44,2,44,42],"489":[45,2,45,55],"497":[45,2,45,56],"505":[47,9,47,10],"513":[47,13,47,14],"521":[47,9,47,14],"529":[47,5,47,15],"537":[48,9,48,10],"545":[48,14,48,15],"553":[49,13,49,17],"561":[49,13,49,17],"569":[49,9,49,18],"577":[51,5,51,17],"585":[51,18,51,25],"593":[51,5,51,26],"601":[51,5,51,27],"609":[43,1,52,2],"617":[43,1,52,2],"625":[54,5,54,11],"633":[54,14,54,20],"641":[55,5,55,17],"649":[55,18,55,25],"657":[55,5,55,26],"665":[55,5,55,27],"673":[9,1,59,2],"681":[9,1,59,2],"689":[9,1,59,2],"697":[9,1,59,2],"705":[9,1,59,2],"713":[9,1,59,2],"721":[19,1,30,2],"729":[9,1,59,2],"737":[32,1,41,2],"745":[9,1,59,2],"753":[43,1,52,2],"761":[9,1,59,2],"769":[64,2,64,8],"777":[64,2,64,10],"785":[64,2,64,11],"793":[66,2,66,8],"801":[66,9,66,10],"809":[66,2,66,11],"817":[66,2,66,12],"825":[65,3,67,2],"833":[1,1,70,1],"841":[5,1,5,32],"849":[1,1,70,1],"857":[9,1,59,2],"865":[1,1,70,1],"873":[5,1,5,32],"881":[5,1,5,32],"889":[23,5,25,6],"897":[26,5,29,6],"905":[19,1,30,2],"913":[19,1,30,2],"921":[36,5,40,6],"929":[32,1,41,2],"937":[32,1,41,2],"945":[48,5,50,6],"953":[43,1,52,2],"961":[43,1,52,2],"969":[54,1,56,2],"977":[9,1,59,2],"985":[9,1,59,2],"993":[1,1,70,1],"1001":[1,1,70,1],"nBranches":14,"originalCodeFileName":"/home/landon/Documents/repos/__active__/quiver_se/bin/example._fninf.js","instrumentedCodeFileName":"/home/landon/Documents/repos/__active__/quiver_se/bin/example._fninf_jalangi_.js","code":"\n// This code has been instrumented by function_info/instrument.js\n\n// bgn decl-prefix (static)\nfunction _Q$xnH(e) { throw e; }\n// end decl-prefix (static)\n\n// bgn entry-point (has-script)\nfunction _Q$ent() {\n\nvar sym__x = \"X:Real\";\nvar sym__y = \"Y:Real\";\n\nvar z = sym__x;\nvar a = false;\n\n\n\nfunction onFirst() {\n\t\"!!MAGIC@js_concolic/src-range=81:259\";\n\t\"!!MAGIC@js_concolic/idents=z:sym__y:a:setImmediate:onSecond\";\n\t\n    if (z === sym__y + 1 && a) {\n        throw 'Stickerbrush? Really?';\n    }\n    if (z < sym__y) {\n        z = z + 2;\n        setImmediate(onSecond)\n    }\n}\n\nfunction onSecond() {\n\t\"!!MAGIC@js_concolic/src-range=261:395\";\n\t\"!!MAGIC@js_concolic/idents=z:sym__y:a:setImmediate:onThird:onFirst\";\n\t\n    if (z === sym__y && !a) {\n        setImmediate(onThird);\n    } else {\n        setImmediate(onFirst);\n    }\n}\n\nfunction onThird() {\n\t\"!!MAGIC@js_concolic/src-range=397:503\";\n\t\"!!MAGIC@js_concolic/idents=z:a:setImmediate:onFirst\";\n\t\n    z = z - 1;\n    if (z != 2) {\n        a = true;\n    }\n    setImmediate(onFirst);\n}\n\nif (sym__x < sym__y) {\n    setImmediate(onFirst);\n}\n\n\n}\n// end entry-point (has-script)\n\n// bgn main-rescue (actual-entry-point)\ntry {\n\t_Q$ent();\n} catch (e) {\n\t_Q$xnH(e);\n}\n// end main-rescue (actual-entry-point)\n\n"};
jalangiLabel5:
    while (true) {
        try {
            J$.Se(833, '/home/landon/Documents/repos/__active__/quiver_se/bin/example._fninf_jalangi_.js', '/home/landon/Documents/repos/__active__/quiver_se/bin/example._fninf.js');
            function _Q$xnH(e) {
                jalangiLabel0:
                    while (true) {
                        try {
                            J$.Fe(33, arguments.callee, this, arguments);
                            arguments = J$.N(41, 'arguments', arguments, 4);
                            e = J$.N(49, 'e', e, 4);
                            throw J$.X1(25, J$.Th(17, J$.R(9, 'e', e, 0)));
                        } catch (J$e) {
                            J$.Ex(873, J$e);
                        } finally {
                            if (J$.Fr(881))
                                continue jalangiLabel0;
                            else
                                return J$.Ra();
                        }
                    }
            }
            function _Q$ent() {
                jalangiLabel4:
                    while (true) {
                        try {
                            J$.Fe(673, arguments.callee, this, arguments);
                            function onFirst() {
                                jalangiLabel1:
                                    while (true) {
                                        try {
                                            J$.Fe(321, arguments.callee, this, arguments);
                                            arguments = J$.N(329, 'arguments', arguments, 4);
                                            J$.X1(161, J$.T(153, "!!MAGIC@js_concolic/src-range=81:259", 21, false));
                                            J$.X1(177, J$.T(169, "!!MAGIC@js_concolic/idents=z:sym__y:a:setImmediate:onSecond", 21, false));
                                            if (J$.X1(889, J$.C(16, J$.C(8, J$.B(18, '===', J$.R(185, 'z', z, 0), J$.B(10, '+', J$.R(193, 'sym__y', sym__y, 0), J$.T(201, 1, 22, false), 0), 0)) ? J$.R(209, 'a', a, 0) : J$._()))) {
                                                throw J$.X1(233, J$.Th(225, J$.T(217, 'Stickerbrush? Really?', 21, false)));
                                            }
                                            if (J$.X1(897, J$.C(24, J$.B(26, '<', J$.R(241, 'z', z, 0), J$.R(249, 'sym__y', sym__y, 0), 0)))) {
                                                J$.X1(281, z = J$.W(273, 'z', J$.B(34, '+', J$.R(257, 'z', z, 0), J$.T(265, 2, 22, false), 0), z, 0));
                                                J$.X1(313, J$.F(305, J$.R(289, 'setImmediate', setImmediate, 2), 0)(J$.R(297, 'onSecond', onSecond, 0)));
                                            }
                                        } catch (J$e) {
                                            J$.Ex(905, J$e);
                                        } finally {
                                            if (J$.Fr(913))
                                                continue jalangiLabel1;
                                            else
                                                return J$.Ra();
                                        }
                                    }
                            }
                            function onSecond() {
                                jalangiLabel2:
                                    while (true) {
                                        try {
                                            J$.Fe(457, arguments.callee, this, arguments);
                                            arguments = J$.N(465, 'arguments', arguments, 4);
                                            J$.X1(345, J$.T(337, "!!MAGIC@js_concolic/src-range=261:395", 21, false));
                                            J$.X1(361, J$.T(353, "!!MAGIC@js_concolic/idents=z:sym__y:a:setImmediate:onThird:onFirst", 21, false));
                                            if (J$.X1(921, J$.C(40, J$.C(32, J$.B(42, '===', J$.R(369, 'z', z, 0), J$.R(377, 'sym__y', sym__y, 0), 0)) ? J$.U(50, '!', J$.R(385, 'a', a, 0)) : J$._()))) {
                                                J$.X1(417, J$.F(409, J$.R(393, 'setImmediate', setImmediate, 2), 0)(J$.R(401, 'onThird', onThird, 0)));
                                            } else {
                                                J$.X1(449, J$.F(441, J$.R(425, 'setImmediate', setImmediate, 2), 0)(J$.R(433, 'onFirst', onFirst, 0)));
                                            }
                                        } catch (J$e) {
                                            J$.Ex(929, J$e);
                                        } finally {
                                            if (J$.Fr(937))
                                                continue jalangiLabel2;
                                            else
                                                return J$.Ra();
                                        }
                                    }
                            }
                            function onThird() {
                                jalangiLabel3:
                                    while (true) {
                                        try {
                                            J$.Fe(609, arguments.callee, this, arguments);
                                            arguments = J$.N(617, 'arguments', arguments, 4);
                                            J$.X1(481, J$.T(473, "!!MAGIC@js_concolic/src-range=397:503", 21, false));
                                            J$.X1(497, J$.T(489, "!!MAGIC@js_concolic/idents=z:a:setImmediate:onFirst", 21, false));
                                            J$.X1(529, z = J$.W(521, 'z', J$.B(58, '-', J$.R(505, 'z', z, 0), J$.T(513, 1, 22, false), 0), z, 0));
                                            if (J$.X1(945, J$.C(48, J$.B(66, '!=', J$.R(537, 'z', z, 0), J$.T(545, 2, 22, false), 0)))) {
                                                J$.X1(569, a = J$.W(561, 'a', J$.T(553, true, 23, false), a, 0));
                                            }
                                            J$.X1(601, J$.F(593, J$.R(577, 'setImmediate', setImmediate, 2), 0)(J$.R(585, 'onFirst', onFirst, 0)));
                                        } catch (J$e) {
                                            J$.Ex(953, J$e);
                                        } finally {
                                            if (J$.Fr(961))
                                                continue jalangiLabel3;
                                            else
                                                return J$.Ra();
                                        }
                                    }
                            }
                            arguments = J$.N(681, 'arguments', arguments, 4);
                            J$.N(689, 'sym__x', sym__x, 0);
                            J$.N(697, 'sym__y', sym__y, 0);
                            J$.N(705, 'z', z, 0);
                            J$.N(713, 'a', a, 0);
                            onFirst = J$.N(729, 'onFirst', J$.T(721, onFirst, 12, false, 321), 0);
                            onSecond = J$.N(745, 'onSecond', J$.T(737, onSecond, 12, false, 457), 0);
                            onThird = J$.N(761, 'onThird', J$.T(753, onThird, 12, false, 609), 0);
                            var sym__x = J$.X1(73, J$.W(65, 'sym__x', J$.T(57, "X:Real", 21, false), sym__x, 1));
                            var sym__y = J$.X1(97, J$.W(89, 'sym__y', J$.T(81, "Y:Real", 21, false), sym__y, 1));
                            var z = J$.X1(121, J$.W(113, 'z', J$.R(105, 'sym__x', sym__x, 0), z, 1));
                            var a = J$.X1(145, J$.W(137, 'a', J$.T(129, false, 23, false), a, 1));
                            if (J$.X1(969, J$.C(56, J$.B(74, '<', J$.R(625, 'sym__x', sym__x, 0), J$.R(633, 'sym__y', sym__y, 0), 0)))) {
                                J$.X1(665, J$.F(657, J$.R(641, 'setImmediate', setImmediate, 2), 0)(J$.R(649, 'onFirst', onFirst, 0)));
                            }
                        } catch (J$e) {
                            J$.Ex(977, J$e);
                        } finally {
                            if (J$.Fr(985))
                                continue jalangiLabel4;
                            else
                                return J$.Ra();
                        }
                    }
            }
            _Q$xnH = J$.N(849, '_Q$xnH', J$.T(841, _Q$xnH, 12, false, 33), 0);
            _Q$ent = J$.N(865, '_Q$ent', J$.T(857, _Q$ent, 12, false, 673), 0);
            try {
                J$.X1(785, J$.F(777, J$.R(769, '_Q$ent', _Q$ent, 1), 0)());
            } catch (e) {
                e = J$.N(825, 'e', e, 1);
                J$.X1(817, J$.F(809, J$.R(793, '_Q$xnH', _Q$xnH, 1), 0)(J$.R(801, 'e', e, 0)));
            }
        } catch (J$e) {
            J$.Ex(993, J$e);
        } finally {
            if (J$.Sr(1001)) {
                J$.L();
                continue jalangiLabel5;
            } else {
                J$.L();
                break jalangiLabel5;
            }
        }
    }
// JALANGI DO NOT INSTRUMENT
